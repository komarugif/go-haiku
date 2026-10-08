// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// Haiku's issetugid always returns 1 ("as long as we're effectively a
// single user system", libbsd): taken as it is, every program would be in
// secure mode, which hides the traceback of a fatal signal. As on AIX, the
// IDs tell instead.

// secureMode is only ever mutated in schedinit, so we don't need to worry about
// synchronization primitives.
var secureMode bool

func initSecureMode() {
	secureMode = !(getuid() == geteuid() && getgid() == getegid())
}

func isSecureMode() bool {
	return secureMode
}
