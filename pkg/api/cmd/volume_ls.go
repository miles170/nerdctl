/*
   Copyright The containerd Authors.

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

package cmd

import (
	"github.com/containerd/nerdctl/pkg/api/types"

	"github.com/spf13/cobra"
)

func VolumeLsFlags(cmd *cobra.Command, opts *types.VolumeLsOptions) {
	flags := cmd.Flags()

	flags.BoolVarP(&opts.Quiet, "quiet", "q", false, "Only display volume names")
	// Alias "-f" is reserved for "--filter"
	flags.StringVar(&opts.Format, "format", "", "Format the output using the given go template")
	flags.BoolVarP(&opts.Size, "size", "s", false, "Display the disk usage of volumes. Can be slow with volumes having loads of directories.")
	cmd.RegisterFlagCompletionFunc("format", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"json", "table", "wide"}, cobra.ShellCompDirectiveNoFileComp
	})
	flags.StringSliceVarP(&opts.Filters, "filter", "f", []string{}, "Filter matches volumes based on given conditions")
}
