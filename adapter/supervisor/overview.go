/*
Package supervisor provides the development kit for running Blacksmith applications
in a distributed environment. The supervisor is used to make sure resources are
accessed by a single instance of the gateway and the scheduler to avoid collision
when listening for events, executing jobs, or running migrations.

Note: The supervisor is part of Blacksmith Enterprise Edition and is not leveraged
when using Blacksmith Standard Edition.
*/
package supervisor
