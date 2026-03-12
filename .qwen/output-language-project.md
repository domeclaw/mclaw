# Output language preference: Thai
<!-- qwen-code:llm-output-language: Thai -->

## Goal
Prefer responding in **Thai** for normal assistant messages and explanations in the context of the MClaw project.

## Keep technical artifacts unchanged
Do **not** translate or rewrite:
- Code blocks, CLI commands, file paths, stack traces, logs, JSON keys, identifiers related to MClaw project
- Exact quoted text from the user (keep quotes verbatim)
- Technical terms, variable names, function names, and configuration keys in the MClaw codebase

## Project-specific context
- MClaw is a lightweight AI assistant with wallet and webhook functionality
- The system has specific files like IDENTITY.md, config.json, and various Go source files
- When discussing project internals, maintain English for technical components while explaining in Thai

## When a conflict exists
If higher-priority instructions (system/developer) require a different behavior, follow them.

## Tool / system outputs
Raw tool/system outputs may contain fixed-format English. Preserve them verbatim, and if needed, add a short **Thai** explanation below.

## Git and project workflows
- Git commit messages must remain in English regardless of the response language
- Project-specific commands and file paths should be preserved in English
- Configuration values and technical parameters should remain in English