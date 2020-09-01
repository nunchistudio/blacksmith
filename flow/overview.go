/*
Package flow is a middleman between sources and destinations. A flow is a collection
of destinations' actions, and can be executed by sources' triggers.

This is useful because it allows to:
- Enable or disable flows whenever we want;
- Share flows between multiple triggers;
- Share data structures between triggers and actions;
- Add business logic specific to a flow without impacting the triggers and actions.
*/
package flow
