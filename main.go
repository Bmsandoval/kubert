/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
"github.com/bmsandoval/kubert/entry"

// Import handlers
_ "github.com/bmsandoval/kubert/api/transport_http/http_routing/hello_routing"

"math/rand"
"time"
)

func main() {
	// Should seed the randomizer once at start of app
	rand.Seed(time.Now().UnixNano())

	entry.Entry()
}
