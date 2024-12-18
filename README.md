## fe3h-enemy-randomizer

A script that randomizes enemy skills and separates them into tiers to allow difficulty-based randomization.

### How does it work

It takes all the skills, classes, weapons, battalions and items in the game and splits them into 6 levels of strength. Depending on how strong I want my enemy to be, there's a certain % a difficulty will be picked for every inventory part.
For instance, a "Level 1" enemy will have a high likelihood of being a generic Thief, having mediocre items, an incomplete skill list, and will generally be easy to mow down, while a late-game boss might end up with 6 Relics, a skill that completely negates Magic damage, and the "King of Liberation" class.

There's an alternate "no restriction" mode where the script generates all it needs regardless of balancing, but I saw one of its outputs and it was _painful_.

![image](https://github.com/user-attachments/assets/6a9324f7-ee62-42a7-829e-9be4354605f4)

Enemies will be edited into the maps using a separate program. Non-boss enemies will have a `n` difficulty assigned to them (`n` being the general map difficulty depending on the story progression) and bosses will have an `n+1` difficulty. So for example, all of Kostas' goons in the tutorial are Level 1, while Kostas himself is Level 2.

Things will get dangerous as I approach Level 4 enemies. I intend on streaming a full Golden Deer run of it on my [YouTube Channel](https://youtube.com/@n_tonio36). See you there!
