# Codebase Review: Proposed Tasks

## 1) Typo fix task
**Task:** Update README wording from "setup cmds" to "setup commands" (and similar shorthand/grammar cleanups in the setup steps).

- **Why:** The README currently uses informal shorthand that reads like a typo and reduces documentation clarity for new contributors.
- **Where:** `README.md` (step 9 currently says "Make Makefile with setup cmds").
- **Acceptance criteria:**
  - Replace obvious shorthand typos (for example, `cmds` -> `commands`).
  - Keep command examples unchanged and executable.

## 2) Bug fix task
**Task:** Fix invalid currency code generation in random test data.

- **Why:** `RandomCurrency()` returns `"YEN"`, but ISO 4217 code for Japanese Yen is `"JPY"`. Using non-standard currency identifiers can break integrations and validation if constraints are added later.
- **Where:** `util/random.go`.
- **Acceptance criteria:**
  - Replace `"YEN"` with `"JPY"`.
  - Add/adjust test coverage to assert only supported ISO-style currency codes are emitted.

## 3) Comment/documentation discrepancy task
**Task:** Reconcile transfer amount DB comment with actual schema constraints.

- **Why:** The schema comment says transfer amounts "can only be positive", but the table has no `CHECK (amount > 0)` constraint, so negatives are currently allowed.
- **Where:** `db/migrate/000001_init_schema.up.sql` (`COMMENT ON COLUMN "transfers"."amount" IS 'Can only be Positive';`).
- **Acceptance criteria:**
  - Either add a DB-level check constraint enforcing positive transfer amounts, **or** update the comment to match actual behavior.
  - Prefer constraint enforcement for data integrity.

## 4) Test improvement task
**Task:** Implement missing transfer tests in `transfer_test.go`.

- **Why:** `db/sqlc/transfer_test.go` currently only contains `package db` and no tests, leaving CRUD/query behavior for transfers unverified.
- **Where:** `db/sqlc/transfer_test.go`.
- **Acceptance criteria:**
  - Add at least create/get/list transfer tests consistent with account/entry test style.
  - Include assertions for IDs, involved account IDs, amounts, and timestamps.
  - Ensure test names are discoverable by `go test`.
