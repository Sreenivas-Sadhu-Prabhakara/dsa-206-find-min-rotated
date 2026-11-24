# Find Minimum in Rotated Sorted Array

> NeetCode 150 · Binary Search · Blind 75 — compare mid against the right end, not the left: that's the comparison that stays unambiguous. Go, stdlib only, tested.

## Problem

A sorted array of distinct values was rotated; find its minimum in O(log n). Companion to Search-in-Rotated (dsa-064) — same terrain, different quarry.

## Approach

The minimum sits just after the "cliff" where the maximum meets it. Compare `nums[mid]` with `nums[hi]`: greater means mid is on the high plateau, so the cliff — and the minimum — lies strictly right (`lo = mid+1`); less-or-equal means mid is past the cliff or the array is unrotated, so the minimum is at mid or left (`hi = mid`, keeping mid alive as a candidate). Why the *right* end and not the left: against `nums[lo]`, the unrotated case and the mid-on-high-plateau case give the same comparison result — ambiguous; against `nums[hi]` the two halves always read differently. The asymmetric shrink (`mid+1` vs `mid`) matches which side provably excludes the answer.

**Complexity:** O(log n) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
