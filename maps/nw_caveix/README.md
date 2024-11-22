# Conjurer campaign map `con01a` script

This script was produced by following the [NS3 migration guide](https://noxworld-dev.github.io/opennox-docs/noxscript/ns4/migrate/index.html).

Most of the code was produced by `noxtools ns decomp`.

It was additionally refactored:
- Cleaned local variables and `goto`s.
- Variables were renamed to make sense.
- Functions and variables grouped to separate files.
- Common function bodies moved to helper functions.
- Code related to shaking was moved to its own struct/class.

This script is a proof-of-concept. It was not yet tested and may work incorrectly.

Specifically, integration between map triggers was not done yet.

However, it is still useful as an example, because all the logic still will work and makes sense.