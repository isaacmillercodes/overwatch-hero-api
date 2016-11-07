TRUNCATE heroes CASCADE;

-- create heroes

INSERT INTO heroes (name, description, role, type, image, difficulty)
VALUES
  ('Genji',
  'Genji flings precise and deadly Shuriken at his targets, and uses his technologically-advanced katana to deflect projectiles or deliver a Swift Strike that cuts down enemies.',
  'Offense',
  'Mobile DPS',
  'https://blzgdapipro-a.akamaihd.net/hero/genji/hero-select-portrait.png',
  3
  ),
  ('McCree',
  'Armed with his Peacekeeper revolver, McCree takes out targets with deadeye precision and dives out of danger with eagle-like speed.',
  'Offense',
  'Team Ult',
  'https://blzgdapipro-a.akamaihd.net/hero/mccree/hero-select-portrait.png',
  2
  ),
  ('Pharah',
  'Soaring through the air in her combat armor, and armed with a launcher that lays down high-explosive rockets, Pharah is a force to be reckoned with.',
  'Offense',
  'Explosive Damage',
  'https://blzgdapipro-a.akamaihd.net/hero/pharah/hero-select-portrait.png',
  1
  ),
  ('Reaper',
  'Hellfire Shotguns, the ghostly ability to become immune to damage, and the power to step between shadows make Reaper one of the deadliest beings on Earth.',
  'Offense',
  'Team Ult',
  'https://blzgdapipro-a.akamaihd.net/hero/reaper/hero-select-portrait.png',
  1
  ),
  ('Soldier: 76',
  'Armed with cutting-edge weaponry, including an experimental pulse rifle that’s capable of firing spirals of high-powered Helix Rockets, Soldier: 76 has the speed and support know-how of a highly trained warrior.',
  'Offense',
  'Self Healing',
  'https://blzgdapipro-a.akamaihd.net/hero/soldier-76/hero-select-portrait.png',
  1
  ),
  ('Sombra',
  'Stealth and debilitating attacks make Sombra a powerful infiltrator. Her hacking can disrupt her enemies, ensuring they’re easier to take out, while her EMP provides the upper hand against multiple foes at once. Sombra’s ability to Translocate and camouflage herself makes her a hard target to pin down.',
  'Offense',
  'Team Assist',
  'https://blzgdapipro-a.akamaihd.net/hero/sombra/hero-select-portrait-d5121256f71c9d7dc7a434ac75be95d99942e8386ba7f8462f3e15d91223854c9b9adde42a3aca70715ab24326a7c27848151e8ab92a259ac7744d7f15a6d91b.png',
  3
  ),
  ('Tracer',
  'Toting twin pulse pistols, energy-based time bombs, and rapid-fire banter, Tracer is able to "blink" through space and rewind her personal timeline as she battles to right wrongs the world over.',
  'Offense',
  'Mobile DPS',
  'https://blzgdapipro-a.akamaihd.net/hero/tracer/hero-select-portrait.png',
  2
  ),
  ('Bastion',
  'Repair protocols and the ability to transform between stationary Assault, mobile Recon and devastating Tank configurations provide Bastion with a high probability of victory.',
  'Defense',
  'Solo Defender',
  'https://blzgdapipro-a.akamaihd.net/hero/bastion/hero-select-portrait.png',
  1
  ),
  ('Hanzo',
  'Hanzo’s versatile arrows can reveal his enemies or fragment to strike multiple targets. He can scale walls to fire his bow from on high, or summon a titanic spirit dragon.',
  'Defense',
  'Sniper',
  'https://blzgdapipro-a.akamaihd.net/hero/hanzo/hero-select-portrait.png',
  3
  ),
  ('Junkrat',
  'Junkrat’s area-denying armaments include a Frag Launcher that lobs bouncing grenades, Concussion Mines that send enemies flying, and Steel Traps that stop foes dead in their tracks.',
  'Defense',
  'Explosive Damage',
  'https://blzgdapipro-a.akamaihd.net/hero/junkrat/hero-select-portrait.png',
  2
  ),
  ('Mei',
  'Mei’s weather-altering devices slow opponents and protect locations. Her Endothermic Blaster unleashes damaging icicles and frost streams, and she can Cryo-Freeze herself to guard against counterattacks, or obstruct the opposing team’s movements with an Ice Wall.',
  'Defense',
  'Solo Defender',
  'https://blzgdapipro-a.akamaihd.net/hero/mei/hero-select-portrait.png',
  3
  ),
  ('Torbjörn',
  'Torbjörn’s extensive arsenal includes a rivet gun and hammer, as well as a personal forge that he can use to build upgradeable turrets and dole out protective armor packs.',
  'Defense',
  'Passive Damage',
  'https://blzgdapipro-a.akamaihd.net/hero/torbjorn/hero-select-portrait.png',
  2
  ),
  ('Widowmaker',
  'Widowmaker equips herself with whatever it takes to eliminate her targets, including mines that dispense poisonous gas, a visor that grants her squad infra-sight, and a powerful sniper rifle that can fire in fully-automatic mode.',
  'Defense',
  'Sniper',
  'https://blzgdapipro-a.akamaihd.net/hero/widowmaker/hero-select-portrait.png',
  2
  ),
  ('D.Va',
  'D.Va’s mech is nimble and powerful—its twin Fusion Cannons blast away with autofire at short range, and she can use its Boosters to barrel over enemies and obstacles, or deflect attacks with her projectile-dismantling Defense Matrix.',
  'Tank',
  'Mobile Tank',
  'https://blzgdapipro-a.akamaihd.net/hero/dva/hero-select-portrait.png',
  2
  ),
  ('Reinhardt',
  'Clad in powered armor and swinging his hammer, Reinhardt leads a rocket-propelled charge across the battleground and defends his squadmates with a massive energy barrier.',
  'Tank',
  'Team Shield',
  'https://blzgdapipro-a.akamaihd.net/hero/reinhardt/hero-select-portrait.png',
  1
  ),
  ('Roadhog',
  'Roadhog uses his signature Chain Hook to pull his enemies close before shredding them with blasts from his Scrap Gun. He’s hardy enough to withstand tremendous damage, and can recover his health with a short breather.',
  'Tank',
  'Self Healing',
  'https://blzgdapipro-a.akamaihd.net/hero/roadhog/hero-select-portrait.png',
  1
  ),
  ('Winston',
  'Winston wields impressive inventions—a jump pack, electricity-blasting Tesla Cannon, portable shield projector and more—with literal gorilla strength.',
  'Tank',
  'Mobile Tank',
  'https://blzgdapipro-a.akamaihd.net/hero/winston/hero-select-portrait.png',
  2
  ),
  ('Zarya',
  'Deploying powerful personal barriers that convert incoming damage into energy for her massive Particle Cannon, Zarya is an invaluable asset on the front lines of any battle.',
  'Tank',
  'Team Assist',
  'https://blzgdapipro-a.akamaihd.net/hero/zarya/hero-select-portrait.png',
  3
  ),
  ('Ana',
  'Ana’s versatile arsenal allows her to affect heroes all over the battlefield. Her Biotic Rifle rounds and Biotic Grenades heal allies and damage or impair enemies; her sidearm tranquilizes key targets, and Nano Boost gives one of her comrades a considerable increase in power.',
  'Support',
  'Debuff Healer',
  'https://blzgdapipro-a.akamaihd.net/hero/ana/hero-select-portrait.png',
  3
  ),
  ('Lúcio',
  'On the battlefield, Lúcio’s cutting-edge Sonic Amplifier buffets enemies with projectiles and knocks foes back with blasts of sound. His songs can both heal his team or boost their movement speed, and he can switch between tracks on the fly.',
  'Support',
  'Solo Healer',
  'https://blzgdapipro-a.akamaihd.net/hero/lucio/hero-select-portrait.png',
  2
  ),
  ('Mercy',
  'Mercy’s Valkyrie Suit helps keep her close to teammates like a guardian angel; healing, resurrecting or strengthening them with the beams emanating from her Caduceus Staff.',
  'Support',
  'Solo Healer',
  'https://blzgdapipro-a.akamaihd.net/hero/mercy/hero-select-portrait.png',
  1
  ),
  ('Symmetra',
  'Symmetra utilizes her light-bending Photon Projector to dispatch adversaries, shield her associates, construct teleportation pads and deploy particle-blasting Sentry Turrets.',
  'Support',
  'Passive Damage',
  'https://blzgdapipro-a.akamaihd.net/hero/symmetra/hero-select-portrait.png',
  2
  ),
  ('Zenyatta',
  'Zenyatta calls upon orbs of harmony and discord to heal his teammates and weaken his opponents, all while pursuing a transcendent state of immunity to damage.',
  'Support',
  'Debuff Healer',
  'https://blzgdapipro-a.akamaihd.net/hero/zenyatta/hero-select-portrait.png',
  3
  );



  TRUNCATE game_types CASCADE;

  -- create game_types

  INSERT INTO game_types (name, has_defense)
  VALUES
    ('Assault', true),
    ('Escort', true),
    ('Hybrid', true),
    ('Control', false);
