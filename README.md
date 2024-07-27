## WHAT

go-prt is a priority scheduler for a group of goRoutines.

## WHY

In certain scenarios, it is required to run (schedule is a better word) a set of task concurrently (not parallelly) in a **priority**
order, since go runtime use a normal run-queue for scheduling the goRoutines.

## Functional Requirements

- To create a library for creating a go scheduler for a group of tasks.
- The scheduler need to interact with the runtime scheduler and handover the tasks to the run-queue in a predefined priority order.