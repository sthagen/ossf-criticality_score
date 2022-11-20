// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package signalio

import (
	"fmt"

	"github.com/ossf/criticality_score/internal/collector/signal"
)

func fieldsFromSignalSets(sets []signal.Set, extra []string) []string {
	var fields []string
	for _, s := range sets {
		if err := signal.ValidateSet(s); err != nil {
			panic(err)
		}
		fields = append(fields, signal.SetFields(s, true)...)
	}
	// Append all the extra fields
	fields = append(fields, extra...)
	return fields
}

func marshalToMap(signals []signal.Set, extra ...Field) (map[string]string, error) {
	values := make(map[string]string)
	for _, s := range signals {
		// Get all of the signal data from the set and serialize it.
		for k, v := range signal.SetAsMap(s, true) {
			if s, err := marshalValue(v); err != nil {
				return nil, fmt.Errorf("failed to write field %s: %w", k, err)
			} else {
				values[k] = s
			}
		}
	}
	for _, f := range extra {
		if s, err := marshalValue(f.Value); err != nil {
			return nil, fmt.Errorf("failed to write field %s: %w", f.Key, err)
		} else {
			values[f.Key] = s
		}
	}
	return values, nil
}