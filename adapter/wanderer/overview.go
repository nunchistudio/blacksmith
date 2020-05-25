/*
Package wanderer provides the development kit for working with database migrations
using a remote mutual exclusion (mutex).

This feature allows teams to run database migrations with no conflicts thanks to
a remote mutex and migration versioning best practices.

A wanderer adapter can be generated using the Blacksmith CLI:

  $ blacksmith generate wanderer

Note: Adapter generation using the Blacksmith CLI is a feature only available in
Blacksmith Enterprise.

Note: The wanderer is part of Blacksmith Enterprise and is not used by adapters
when using the open-source version of Blacksmith.
*/
package wanderer
