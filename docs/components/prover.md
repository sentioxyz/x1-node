# Component: Prover

NOTE: The Prover is not considered part of the XAGON Node and all issues and suggestions should be sent to the [Prover repo](https://github.com/okx/Xagon-prover/).

## XAGON Prover:

The XAGON Prover image hosts different components, *Merkle Tree*, *Executor* and finally the actual *Prover*.

## Hard dependencies:

- [Aggregator](./aggregator.md)

## Running:

The preferred way to run the XAGON Prover component is via Docker and Docker Compose.

```bash
docker pull hermeznetwork/Xagon-prover
```

To orchestrate multiple deployments of the different XAGON Node components, a `docker-compose.yaml` file for Docker Compose can be used:

```yaml
  Xagon-prover:
    container_name: Xagon-prover
    image: Xagon-prover
    volumes:
      - ./prover-config.json:/usr/src/app/config.json
    command: >
      zkProver -c /usr/src/app/config.json
```

The `prover-config.json` file contents will vary depending on your use case, the main document explains different values to be changed to achieve different behaviors.

### Ports:

- `50051`: Prover
- `50061`: Merkle Tree
- `50071`: Executor