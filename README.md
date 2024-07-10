# Article Construct Demo

[![GoTemplate](https://img.shields.io/badge/go/template-black?logo=go)](https://github.com/SchwarzIT/go-template)

Finally, it's time for the product to come to life! Similar to the creation of a render tree in web development, the product's article construct structure and details are used to generate the final user experience. This process involves translating the Article Construct into the appropriate format for your chosen platform (e.g., code for a software application, design mockups for a physical product).

The project uses `make` to make your life easier. If you're not familiar with Makefiles you can take a look at [this quickstart guide](https://makefiletutorial.com).

Whenever you need help regarding the available actions, just use the following command.

```bash
make help
```

## Setup

To get your setup up and running the only thing you have to do is

```bash
make all
```

This will initialize a git repo, download the dependencies in the latest versions and install all needed tools.
If needed code generation will be triggered in this target as well.

## Test & lint

Run linting

```bash
make lint
```

Run tests

```bash
make test
```
