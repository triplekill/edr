package artifacts

// This allows to run an artifact as a plugin.
import (
	"context"
	"fmt"
	"strings"

	actions_proto "www.velocidex.com/golang/velociraptor/actions/proto"
	artifacts_proto "www.velocidex.com/golang/velociraptor/artifacts/proto"
	"www.velocidex.com/golang/velociraptor/utils"
	"www.velocidex.com/golang/vfilter"
)

type ArtifactRepositoryPlugin struct {
	repository *Repository
	children   map[string]vfilter.PluginGeneratorInterface
	prefix     []string
	leaf       *artifacts_proto.Artifact
}

func (self *ArtifactRepositoryPlugin) Print() {
	var children []string
	for k := range self.children {
		children = append(children, k)
	}
	fmt.Printf("prefix '%v', Children %v, Leaf %v\n",
		self.prefix, children, self.leaf != nil)
	for _, v := range self.children {
		v.(*ArtifactRepositoryPlugin).Print()
	}
}

// Define vfilter.PluginGeneratorInterface
func (self *ArtifactRepositoryPlugin) Call(
	ctx context.Context, scope *vfilter.Scope,
	args *vfilter.Dict) <-chan vfilter.Row {
	output_chan := make(chan vfilter.Row)

	if self.leaf == nil {
		scope.Log("Artifact %s not found", strings.Join(self.prefix, "."))
		close(output_chan)
		return output_chan
	}

	request := &actions_proto.VQLCollectorArgs{}
	err := Compile(self.leaf, request)
	if err != nil {
		scope.Log("Artifact %s invalid: %s",
			strings.Join(self.prefix, "."), err.Error())
		close(output_chan)
		return output_chan
	}

	go func() {
		defer close(output_chan)

		// We create a child scope for evaluating the artifact.
		env := vfilter.NewDict()
		for _, request_env := range request.Env {
			env.Set(request_env.Key, request_env.Value)
		}

		// Allow the args to override the artifact defaults.
		for k, v := range *args.ToDict() {
			env.Set(k, v)
		}

		child_scope := scope.Copy().AppendVars(env)
		for _, query := range request.Query {
			vql, err := vfilter.Parse(query.VQL)
			if err != nil {
				scope.Log("Artifact %s invalid: %s",
					strings.Join(self.prefix, "."),
					err.Error())
				return
			}

			child_chan := vql.Eval(ctx, child_scope)
			for {
				row, ok := <-child_chan
				// This query is done - do the
				// next one.
				if !ok {
					break
				}
				output_chan <- row
			}
		}

	}()
	return output_chan
}

func (self *ArtifactRepositoryPlugin) Name() string {
	return strings.Join(self.prefix, ".")
}

func (self *ArtifactRepositoryPlugin) Info(
	scope *vfilter.Scope, type_map *vfilter.TypeMap) *vfilter.PluginInfo {
	return &vfilter.PluginInfo{
		Name: self.Name(),
		Doc:  "A pseudo plugin for accessing the artifacts repository from VQL.",
	}
}

// Define Associative protocol.
type _ArtifactRepositoryPluginAssociativeProtocol struct{}

func _getArtifactRepositoryPlugin(a vfilter.Any) *ArtifactRepositoryPlugin {
	switch t := a.(type) {
	case ArtifactRepositoryPlugin:
		return &t

	case *ArtifactRepositoryPlugin:
		return t

	default:
		return nil
	}
}

func (self _ArtifactRepositoryPluginAssociativeProtocol) Applicable(
	a vfilter.Any, b vfilter.Any) bool {
	if _getArtifactRepositoryPlugin(a) == nil {
		return false
	}

	switch b.(type) {
	case string:
		break
	default:
		return false
	}

	return true
}

func (self _ArtifactRepositoryPluginAssociativeProtocol) GetMembers(
	scope *vfilter.Scope, a vfilter.Any) []string {
	var result []string

	value := _getArtifactRepositoryPlugin(a)
	if value != nil {
		for k := range value.children {
			result = append(result, k)
		}
	}
	return result
}

func (self _ArtifactRepositoryPluginAssociativeProtocol) Associative(
	scope *vfilter.Scope, a vfilter.Any, b vfilter.Any) (vfilter.Any, bool) {
	value := _getArtifactRepositoryPlugin(a)
	if value == nil {
		return nil, false
	}

	key, _ := b.(string)
	child, pres := value.children[key]
	return child, pres
}

func NewArtifactRepositoryPlugin(
	repository *Repository, prefix []string) vfilter.PluginGeneratorInterface {
	result := &ArtifactRepositoryPlugin{
		repository: repository,
		children:   make(map[string]vfilter.PluginGeneratorInterface),
		prefix:     prefix,
	}

	for _, name := range repository.List() {
		components := strings.Split(name, ".")
		if len(components) < len(prefix) ||
			!utils.SlicesEqual(components[:len(prefix)], prefix) {
			continue
		}

		components = components[len(prefix):]

		// We are at a leaf node.
		if len(components) == 0 {
			artifact, _ := repository.Get(name)
			result.leaf = artifact
			return result
		}

		_, pres := result.children[components[0]]
		if !pres {
			result.children[components[0]] = NewArtifactRepositoryPlugin(
				repository, append(prefix, components[0]))
		}
	}

	return result
}
