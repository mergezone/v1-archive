# TODO

Remaining audit remediation items from the consistency-first design audit.

## Priority 1 — Small, isolated fixes

- [x] **L6** Document `!important` exceptions (diff shadow DOM, pane resize) as intentional in PROGRESS.md
- [x] **L4** Drop singular `data-pull-request-scope` attribute now that plural payload exists
- [x] **D6** Remove `pane_content.templ` / `PullRequestDetail` fragment; compose context pane from smaller components instead

## Priority 2 — Component consolidation

- [x] **D4** Replace custom mobile dock with DaisyUI `dock` component
- [x] **D2** Parameterize sort controls in shared `FilterSectionHeading` / `CollapsibleSectionHeading`
- [x] **L2** Fold `scopes.templ` and `contributors.templ` into `sidebar_card.templ`
- [x] **U4** Consolidate mobile filter panel and desktop sidebar into a single handler

## Priority 3 — Data flow

- [x] **L3** Flip context pane to `data-*`-first population, drop JSON seed where redundant
- [x] **U1** Add `?pr=` to HTMX push-url so shared PR survives filter navigation

## Priority 4 — CSS rework (major effort, incremental)

- [ ] **D1** Audit `app.css` class-by-class against Tailwind/DaisyUI; convert simple ones
- [ ] **L1** Further consolidate responsive and density media queries
- [x] Group `@media (max-width: 63.999rem)` backdrop rules into a single block
- [x] Self-host Manrope or document the Google Fonts dependency

## Follow-ups

- Screenshot-based visual regression tests
- Consolidate deprecated `scopes` / `contributors` templates
- Diff rendering still depends on `@pierre/diffs`
EOF