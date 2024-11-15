# Pandemonium API

## Project 1 - Documentation System

### Obejctives
Build a documentation system that allows voice, text and video input. This system should act as a knowledge base and should be accessible via the different inputs.

Example implementation: Saying "Hey Google, start documentation" which will trigger a new recording by a microphone or camera and record the user. Saying "stop documentation" will stop the process, and once stopped, the voice memo or video should be transcribed to text and JSON data which includes the date,time and location of the documentation, set a general description of not more than 20 words before proceeding to store itself in my knowledge base.

### Implementation Flow
1. Setup NextCloud to sync between devices - Apps like Obsidian
2. Setup API endpoint to create files on nextcloud path