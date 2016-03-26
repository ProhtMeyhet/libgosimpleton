Current
=======

0.2beta
=======

- cleanup + bugfixes
- move simple things to simpleton/*
- logging is now done in a separate goroutine
- logging: add tests. coverage: 87.4%
- refactor of logging api
+ cache (ObservingFileCache needs rethinking)

0.1
===
initial release



libgosimpleton-next
===================

0.2
===
+ simplee

0.3
===
+ add config for SQL in credentialstool
- more tests for credentials
- add iotool

N.x
===
- move crediantialstool to own package
- refactor cache to CachedReader for io.Reader
- add Liltemplate
- add parallel
