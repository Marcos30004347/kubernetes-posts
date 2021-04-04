package server

import (
	"fmt"
	"io"
	"net"

	"github.com/spf13/cobra"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	utilfeature "k8s.io/apiserver/pkg/util/feature"

	"github.com/Marcos30004347/kubernetes-posts/pkg/admission/initializer"
	"github.com/Marcos30004347/kubernetes-posts/pkg/admission/plugin/foobar"
	"github.com/Marcos30004347/kubernetes-posts/pkg/apis/baz/v1alpha1"
	"github.com/Marcos30004347/kubernetes-posts/pkg/apiserver"

	clientset "github.com/Marcos30004347/kubernetes-posts/pkg/generated/clientset/versioned"

	informers "github.com/Marcos30004347/kubernetes-posts/pkg/generated/informers/externalversions"

	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/features"
	genericapiserver "k8s.io/apiserver/pkg/server"
	serveroptions "k8s.io/apiserver/pkg/server/options"
)

type CustomServerOptions struct {
	RecommendedOptions    *serveroptions.RecommendedOptions
	SharedInformerFactory informers.SharedInformerFactory
	StdOut                io.Writer
	StdErr                io.Writer
}

func NewCustomServerOptions(out, errOut io.Writer) *CustomServerOptions {
	// Instantiate the RecommendedOptions
	o := &CustomServerOptions{
		RecommendedOptions: serveroptions.NewRecommendedOptions(
			"/registry/baz",
			apiserver.Codecs.LegacyCodec(v1alpha1.SchemeGroupVersion),
		),

		StdOut: out,
		StdErr: errOut,
	}

	o.RecommendedOptions.Etcd.StorageConfig.EncodeVersioner = runtime.NewMultiGroupVersioner(v1alpha1.SchemeGroupVersion, schema.GroupKind{Group: v1alpha1.GroupName})

	return o
}

// NewCommandStartCustomServer provides a CLI handler for 'start master' command
// with a default CustomServerOptions.
func NewCommandStartCustomServer(
	defaults *CustomServerOptions,
	stopCh <-chan struct{},
) *cobra.Command {
	o := *defaults
	cmd := &cobra.Command{
		Short: "Launch a custom API server",
		Long:  "Launch a custom API server",
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			if err := o.Run(stopCh); err != nil {
				return err
			}
			return nil
		},
	}

	flags := cmd.Flags()

	o.RecommendedOptions.AddFlags(flags)

	utilfeature.DefaultMutableFeatureGate.AddFlag(flags)

	return cmd
}

func (o *CustomServerOptions) Config() (*apiserver.Config, error) {
	// Tell the recomended options to create a signed certificate if user did not specify it in the flag options
	if err := o.RecommendedOptions.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, []net.IP{net.ParseIP("127.0.0.1")}); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}

	o.RecommendedOptions.Etcd.StorageConfig.Paging = utilfeature.DefaultFeatureGate.Enabled(features.APIListChunking)

	// Here is the setup for the client and informers
	o.RecommendedOptions.ExtraAdmissionInitializers = func(c *genericapiserver.RecommendedConfig) ([]admission.PluginInitializer, error) {
		client, err := clientset.NewForConfig(c.LoopbackClientConfig)
		if err != nil {
			return nil, err
		}
		informerFactory := informers.NewSharedInformerFactory(client, c.LoopbackClientConfig.Timeout)
		o.SharedInformerFactory = informerFactory
		return []admission.PluginInitializer{initializer.New(informerFactory)}, nil
	}

	// Instantiate the default recommended configuration
	serverConfig := genericapiserver.NewRecommendedConfig(apiserver.Codecs)

	// Change the default according to flags and other customized options
	err := o.RecommendedOptions.ApplyTo(serverConfig)

	if err != nil {
		return nil, err
	}

	config := &apiserver.Config{
		GenericConfig: serverConfig,
		ExtraConfig:   apiserver.ExtraConfig{},
	}

	return config, nil
}

func (o CustomServerOptions) Run(stopCh <-chan struct{}) error {
	config, err := o.Config()
	if err != nil {
		return err
	}

	server, err := config.Complete().New()
	if err != nil {
		return err
	}

	server.GenericAPIServer.AddPostStartHook("start-baz-api-informers", func(context genericapiserver.PostStartHookContext) error {
		config.GenericConfig.SharedInformerFactory.Start(context.StopCh)
		o.SharedInformerFactory.Start(context.StopCh)
		return nil
	})

	return server.GenericAPIServer.PrepareRun().Run(stopCh)
}

func (o CustomServerOptions) Validate() error {
	errors := []error{}
	errors = append(errors, o.RecommendedOptions.Validate()...)
	return utilerrors.NewAggregate(errors)
}

func (o *CustomServerOptions) Complete() error {
	// register admission plugins
	foobar.Register(o.RecommendedOptions.Admission.Plugins)

	// add admisison plugins to the RecommendedPluginOrder
	o.RecommendedOptions.Admission.RecommendedPluginOrder = append(o.RecommendedOptions.Admission.RecommendedPluginOrder, "FooBar")

	return nil
}
