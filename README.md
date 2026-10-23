# lock-file-consumer

Consumes both test repositories at `v3.0.0`, locked with `gh actions-lock`:

- [Exypte/lock-file-action](https://github.com/Exypte/lock-file-action) — composite action, used as a step
- [Exypte/lock-file-workflow](https://github.com/Exypte/lock-file-workflow) — reusable workflow, called as a job

The lockfile is `.github/workflows/actions.lock`; regenerate it with `gh actions-lock`.
