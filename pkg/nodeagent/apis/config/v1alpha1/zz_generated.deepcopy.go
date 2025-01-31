//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v3 "github.com/Masterminds/semver/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	configv1alpha1 "k8s.io/component-base/config/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServer) DeepCopyInto(out *APIServer) {
	*out = *in
	if in.CABundle != nil {
		in, out := &in.CABundle, &out.CABundle
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServer.
func (in *APIServer) DeepCopy() *APIServer {
	if in == nil {
		return nil
	}
	out := new(APIServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapConfiguration) DeepCopyInto(out *BootstrapConfiguration) {
	*out = *in
	if in.KubeletDataVolumeSize != nil {
		in, out := &in.KubeletDataVolumeSize, &out.KubeletDataVolumeSize
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapConfiguration.
func (in *BootstrapConfiguration) DeepCopy() *BootstrapConfiguration {
	if in == nil {
		return nil
	}
	out := new(BootstrapConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerConfiguration) DeepCopyInto(out *ControllerConfiguration) {
	*out = *in
	in.OperatingSystemConfig.DeepCopyInto(&out.OperatingSystemConfig)
	out.Token = in.Token
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerConfiguration.
func (in *ControllerConfiguration) DeepCopy() *ControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAgentConfiguration) DeepCopyInto(out *NodeAgentConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ClientConnection = in.ClientConnection
	in.Server.DeepCopyInto(&out.Server)
	if in.Debugging != nil {
		in, out := &in.Debugging, &out.Debugging
		*out = new(configv1alpha1.DebuggingConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.APIServer.DeepCopyInto(&out.APIServer)
	if in.Bootstrap != nil {
		in, out := &in.Bootstrap, &out.Bootstrap
		*out = new(BootstrapConfiguration)
		(*in).DeepCopyInto(*out)
	}
	in.Controllers.DeepCopyInto(&out.Controllers)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAgentConfiguration.
func (in *NodeAgentConfiguration) DeepCopy() *NodeAgentConfiguration {
	if in == nil {
		return nil
	}
	out := new(NodeAgentConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeAgentConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatingSystemConfigControllerConfig) DeepCopyInto(out *OperatingSystemConfigControllerConfig) {
	*out = *in
	if in.SyncPeriod != nil {
		in, out := &in.SyncPeriod, &out.SyncPeriod
		*out = new(v1.Duration)
		**out = **in
	}
	if in.SyncJitterPeriod != nil {
		in, out := &in.SyncJitterPeriod, &out.SyncJitterPeriod
		*out = new(v1.Duration)
		**out = **in
	}
	if in.KubernetesVersion != nil {
		in, out := &in.KubernetesVersion, &out.KubernetesVersion
		*out = new(v3.Version)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatingSystemConfigControllerConfig.
func (in *OperatingSystemConfigControllerConfig) DeepCopy() *OperatingSystemConfigControllerConfig {
	if in == nil {
		return nil
	}
	out := new(OperatingSystemConfigControllerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Server) DeepCopyInto(out *Server) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server.
func (in *Server) DeepCopy() *Server {
	if in == nil {
		return nil
	}
	out := new(Server)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerConfiguration) DeepCopyInto(out *ServerConfiguration) {
	*out = *in
	if in.HealthProbes != nil {
		in, out := &in.HealthProbes, &out.HealthProbes
		*out = new(Server)
		**out = **in
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new(Server)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerConfiguration.
func (in *ServerConfiguration) DeepCopy() *ServerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ServerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenControllerConfig) DeepCopyInto(out *TokenControllerConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenControllerConfig.
func (in *TokenControllerConfig) DeepCopy() *TokenControllerConfig {
	if in == nil {
		return nil
	}
	out := new(TokenControllerConfig)
	in.DeepCopyInto(out)
	return out
}
