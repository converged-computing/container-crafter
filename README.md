# Container Crafter

```console
   _   _   _   _   _   _   _   _   _     _   _   _   _   _   _   _  
  / \ / \ / \ / \ / \ / \ / \ / \ / \   / \ / \ / \ / \ / \ / \ / \ 
 ( c | o | n | t | a | i | n | e | r ) ( c | r | a | f | t | e | r )
  \_/ \_/ \_/ \_/ \_/ \_/ \_/ \_/ \_/   \_/ \_/ \_/ \_/ \_/ \_/ \_/  

```

I want a tool that can programatically generate sets of containers with the following features:

- Control the number of layers and total size of image
- Read an experiment from a configuration file
- Build to a maximum size, size, or size distribution per layer

## Goals

The goals are to be able to do a controlled experiment that varies the size of the containers, both total and number of layers. We would want to be able to answer the following questions:

1. Given the equivalent total size, does it take longer (and thus more costly) to pull many smaller layers, or few larger ones?
2. What happens to this pattern as the pulls are scaled across many nodes?

The experiment I had in mind (automation, configs, etc) is going to be under the [container-chonks](https://github.com/converged-computing/container-chonks/tree/main/experiments/pulling) experiment set, where I'm looking at containers across the ecosystem and assessing pull times.

## Usage

### Parameter Space

For the study we are doing, we are interested in these maximum sizes:

- 53,702,097 bytes  (25th percentile)
- 132,399,102 bytes  (50th percentile)
- 392,602,448 bytes  (75th percentile)
- 19,039,736,629 bytes (100th percentile)

These are percentiles of total sizes from the [Dockerfile database](https://github.com/converged-computing/container-chonks/tree/main/experiments/dockerfile), which we consider a reasonable sample of the ecosystem. For the number of layers:

- 6 (25th percentile)
- 9 (50th percentile)
- 14 (75th percentile)
- 153 (100th percentile) 

Note that Docker does not allow you to build over 127 layers, although it will technically work for other container runtimes. We are setting a limit at 127 to mirror what the average user would have access to. 

### Matrix Generation

We and then are going to generate our builds from a configuration file [examples/study.yaml](examples/study.yaml).
The matrix will be represented in a list of images, where each image has a particular total size and number of layers, and the
layer size for each is calculated based on that.

```bash
# Build all images
./bin/container-crafter create --config ./example/study.yaml 

# Push (they have common URI)
docker push ghcr.io/converged-computing/container-chonks --all-tags
```

After push, depending on the registry you might need to make the images public. If you need to cleanup locally:

```bash
docker rmi $(docker images --filter=reference="*ghcr.io/converged-computing/container-chonks*" -q)
```

The images have layers that are arbitrary content written to a random file name, so they are unique.
This means if you are doing a pulling study, the cache won't be used (which is the goal).
Don't forget to make the repository public, if appropriate.

## License

HPCIC DevTools is distributed under the terms of the MIT license.
All new contributions must be made under this license.

See [LICENSE](https://github.com/converged-computing/cloud-select/blob/main/LICENSE),
[COPYRIGHT](https://github.com/converged-computing/cloud-select/blob/main/COPYRIGHT), and
[NOTICE](https://github.com/converged-computing/cloud-select/blob/main/NOTICE) for details.

SPDX-License-Identifier: (MIT)

LLNL-CODE- 842614