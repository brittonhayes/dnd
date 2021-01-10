package mocks

const (
	MonstersFindAbolethMock Mock = `{"index":"aboleth","name":"Aboleth","size":"Large","type":"aberration","subtype":null,"alignment":"lawful evil","armor_class":17,"hit_points":135,"hit_dice":"18d10","speed":{"walk":"10 ft.","swim":"40 ft."},"strength":21,"dexterity":9,"constitution":15,"intelligence":18,"wisdom":15,"charisma":18,"proficiencies":[{"value":6,"proficiency":{"index":"saving-throw-con","name":"Saving Throw: CON","url":"/api/proficiencies/saving-throw-con"}},{"value":8,"proficiency":{"index":"saving-throw-int","name":"Saving Throw: INT","url":"/api/proficiencies/saving-throw-int"}},{"value":6,"proficiency":{"index":"saving-throw-wis","name":"Saving Throw: WIS","url":"/api/proficiencies/saving-throw-wis"}},{"value":12,"proficiency":{"index":"skill-history","name":"Skill: History","url":"/api/proficiencies/skill-history"}},{"value":10,"proficiency":{"index":"skill-perception","name":"Skill: Perception","url":"/api/proficiencies/skill-perception"}}],"damage_vulnerabilities":[],"damage_resistances":[],"damage_immunities":[],"condition_immunities":[],"senses":{"darkvision":"120 ft.","passive_perception":20},"languages":"Deep Speech, telepathy 120 ft.","challenge_rating":10,"xp":5900,"special_abilities":[{"name":"Amphibious","desc":"The aboleth can breathe air and water."},{"name":"Mucous Cloud","desc":"While underwater, the aboleth is surrounded by transformative mucus. A creature that touches the aboleth or that hits it with a melee attack while within 5 ft. of it must make a DC 14 Constitution saving throw. On a failure, the creature is diseased for 1d4 hours. The diseased creature can breathe only underwater.","dc":{"dc_type":{"index":"con","name":"CON","url":"/api/ability-scores/con"},"dc_value":14,"success_type":"none"}},{"name":"Probing Telepathy","desc":"If a creature communicates telepathically with the aboleth, the aboleth learns the creature's greatest desires if the aboleth can see the creature."}],"actions":[{"name":"Multiattack","desc":"The aboleth makes three tentacle attacks.","options":{"choose":1,"from":[[{"name":"Tentacle","count":3,"type":"melee"}]]},"damage":[]},{"name":"Tentacle","desc":"Melee Weapon Attack: +9 to hit, reach 10 ft., one target. Hit: 12 (2d6 + 5) bludgeoning damage. If the target is a creature, it must succeed on a DC 14 Constitution saving throw or become diseased. The disease has no effect for 1 minute and can be removed by any magic that cures disease. After 1 minute, the diseased creature's skin becomes translucent and slimy, the creature can't regain hit points unless it is underwater, and the disease can be removed only by heal or another disease-curing spell of 6th level or higher. When the creature is outside a body of water, it takes 6 (1d12) acid damage every 10 minutes unless moisture is applied to the skin before 10 minutes have passed.","attack_bonus":9,"dc":{"dc_type":{"index":"con","name":"CON","url":"/api/ability-scores/con"},"dc_value":14,"success_type":"none"},"damage":[{"damage_type":{"index":"acid","name":"Acid","url":"/api/damage-types/acid"},"damage_dice":"2d6+5"}]},{"name":"Tail","desc":"Melee Weapon Attack: +9 to hit, reach 10 ft. one target. Hit: 15 (3d6 + 5) bludgeoning damage.","attack_bonus":9,"damage":[{"damage_type":{"index":"bludgeoning","name":"Bludgeoning","url":"/api/damage-types/bludgeoning"},"damage_dice":"3d6+5"}]},{"name":"Enslave","desc":"<<PRESENCE>>","attack_bonus":0,"damage":[{"damage_type":{"index":"psychic","name":"Psychic","url":"/api/damage-types/psychic"},"damage_dice":"3d6"}]}],"url":"/api/monsters/aboleth"}`
	MonstersListMock        Mock = `{"count":332,"results": "<<PRESENCE>>"}`
)