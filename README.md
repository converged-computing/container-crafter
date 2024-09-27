# Container Crafter

I want a tool that can programatically generate sets of containers with the following features:

- Control the number of layers, and shared layers between the set
- Build to a maximum size, size, or size distribution per layer

## Goals

The goals are to be able to do a controlled experiment that varies the size of the containers, both total and number of layers. We would want to be able to answer the following questions:

1. Given the equivalent total size, does it take longer (and thus more costly) to pull many smaller layers, or few larger ones?
2. What happens to this pattern as the pulls are scaled across many nodes?

## License

HPCIC DevTools is distributed under the terms of the MIT license.
All new contributions must be made under this license.

See [LICENSE](https://github.com/converged-computing/cloud-select/blob/main/LICENSE),
[COPYRIGHT](https://github.com/converged-computing/cloud-select/blob/main/COPYRIGHT), and
[NOTICE](https://github.com/converged-computing/cloud-select/blob/main/NOTICE) for details.

SPDX-License-Identifier: (MIT)

LLNL-CODE- 842614