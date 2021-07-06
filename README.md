# elasticsearch-golang-project

## Usage

### Development

Init project with
```
./init.sh
```

Start project with
```
docker-compose up -d
```

Check logs with
```
docker-compose logs -f
```

## Conventions

### Commit convention

A commit convention based on Angular commit message guideline. You can find it [here](https://gist.github.com/brianclements/841ea7bffdb01346392c)

To make it quick every commit should look like this :

`<type>(<scope>): <subject>`

A git hook is defined in the .githooks directory to respect a certain commit pattern.

Type should be either feat|fix|doc|add|update|delete.

The developer is free to chose scope and subject.

### Branching convention

We will be using [Gitflow](https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow) branching workflow.

A git hook is also defined in the .githooks directory to respect a certain branch naming pattern.

You should call your branch feature|bugfix|improvement|release|hotfix|support/branch-subject