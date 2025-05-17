# 📘 Phase 1 Documentation: Backend Foundations

## 🧭 Overview

Phase 1 of buddy.me focused on building the **system of record**—the core backend foundation modeling users, roadmaps, and checkpoints. The priority was simplicity: to get records that work and are easily testable. Modifications can be applied at a later statge. Phase 1 will also lay the groundwork for RHIA and user-facing logic.

![backend sample](new.gif)

---

## 🏗️ Architecture Summary

| Layer     | Tools/Tech         | Notes                                                                 |
| --------- | ------------------ | --------------------------------------------------------------------- |
| Transport | gRPC               | Contract-based APIs with strict versioning via Protobuf               |
| Database  | PostgreSQL         | Declarative schema management and migrations with `sqlc`              |
| Queries   | SQLC               | Typed, compiled queries; Go bindings separate from SQL logic          |
| Testing   | pgx.Tx             | Table-driven tests with rollback isolation for reliable state cleanup |
| Logging   | Zerolog            | Fast, structured logging                                              |
| Contracts | Protobuf v1        | Defined in `proto/*/v1`, backward compatible                          |
| Dev UX    | grpcurl + Makefile | Reproducible gRPC calls, format/lint/test runners                     |

---

## 📦 Modules

### 1. Users

**Entity:**

```protobuf
message User {
  string id = 1;
  string username = 2;
  string name = 3;
  string email = 4;
  google.protobuf.Timestamp created_at = 5;
}
```

**Endpoints:**

* `CreateUser`
* `GetUser`
* `ListUsers`
* *(UpdateUser and DeleteUser: deferred to Phase 2)*

**Testing:**

* Full test coverage using `pgx.Tx` isolation
* Table-driven format with multiple scenarios
* grpcurl script validation (via `make & grpcurl`)

---

### 2. Roadmaps

**Entity:**

```protobuf
message Roadmap {
  string id = 1;
  string title = 2;
  string description = 3;
  string type = 4;
  string status = 5;
  google.protobuf.Timestamp created_at = 6;
}
```

**Endpoints:**

* `CreateRoadmap`
* `GetRoadmap`
* `ListRoadmaps`
* `DeleteRoadmap`

**Design Notes:**

* System-authored (admin-only)
* Non-editable by users
* Acts as curriculum scaffolding for shared progress

---

### 3. Checkpoints

**Entity:**

```protobuf
message Checkpoint {
  string id = 1;
  string roadmap_id = 2;
  string title = 3;
  int32 sequence = 4;
  string summary = 5;
  google.protobuf.StringValue next_step = 6;
}
```

**Endpoints:**

* `CreateCheckpoint`
* `GetCheckpoint`
* `ListCheckpoints`
* `DeleteCheckpoint`

**Design Notes:**

* Tied to roadmaps; each checkpoint is a sequence step
* Not modifiable or deletable by users
* Users can skip checkpoints, like optional modules

---

### 4. Events *(Stubbed for Phase 2)*

* Placeholder schema for event logging (e.g., completed checkpoint)
* Will be consumed by RHIA in Phase 2

---

## 🔍 Testing Strategy

* **pgx.Tx Rollback:** Each test wrapped in a transaction rollback for isolation
* **Timeout Contexts:** Prevents long-running tests and improves feedback loops
* **Table-Driven Style:** Reusable helpers, clean assertions, edge case enumeration
* **grpcurl Verification:** Ensures interface contracts hold under manual and scripted invocations

---

## 🖼️ Supplemental Docs & Artifacts

To be placed under `docs/`:

* [`docs/users/README.md`](./docs/users/README.md) => (protobuf + handler + test breakdown)
* `docs/roadmaps/README.md`
* `docs/checkpoints/README.md`
* ERD diagram
* grpcurl screenshots / CLI examples
* Test output + code snippets
* Visual diagrams of entity relations and flows

---

## 🧾 Future Developer Articles (Draft Concepts)

### 1. **Designing a System of Record in Go + PostgreSQL**

* Why SQLC over ORMs
* `pgx.Tx` testing methodology
* How versioning and protobuf contracts provide stability

### 2. **Why Roadmaps Aren’t Editable: Designing for Progress, not Preference**

* Gamified structure vs. user agency
* Avoiding overchoice: product philosophy
* Like Duolingo, but for growth

### 3. **Building Intent-first APIs with Protobuf and gRPC**

* Designing for precision and clarity
* Thinking from usage, not just data
* Making endpoints reflect product intentions

---

## ✅ Phase 1 Completion Criteria (All Done)

| Task                              | Status |
| --------------------------------- | ------ |
| Users CRUD via gRPC               | ✅      |
| Roadmaps CRUD via gRPC (admin)    | ✅      |
| Checkpoints CRUD via gRPC (admin) | ✅      |
| pgx.Tx test coverage              | ✅      |
| grpcurl interface tests           | ✅      |
| Logging & migrations              | ✅      |
| Structured Makefile               | ✅      |
| Docs draft for Phase 1            | ✅      |
