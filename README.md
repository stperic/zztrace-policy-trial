# zztrace policy trial

Synthetic qualification fixture for zztrace T3b. It contains no customer data.

The approval trial changes the accepted limit from 50 to 40, updates its code and test, and rechecks its links. The required `zztrace/policy` status must fail until `stperic` approves the exact PR head. A later push needs a new review.

The denial trial changes a function covered by a frozen lock. The required status must fail even if someone approves the PR.

The workflow downloads a checker built from the zztrace source and verifies its SHA-256 before running it. Candidate application code is parsed, never executed, in that job. Local fixture tests are run before publication.

This is a controlled test repository. Its workflow and repository settings are managed by the project owner. The workflow does not demonstrate separation from an administrator who can change that configuration. The main zztrace acceptance report records the exact source, binary and PR results.
