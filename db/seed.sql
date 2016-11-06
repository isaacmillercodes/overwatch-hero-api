TRUNCATE heroes CASCADE;

-- create heroes

INSERT INTO heroes (name, description, role, type, difficulty)
VALUES
  ('Genji',
  'Genji flings precise and deadly Shuriken at his targets, and uses his technologically-advanced katana to deflect projectiles or deliver a Swift Strike that cuts down enemies.',
  'Offense',
  'Mobile DPS',
  3
  ),
  ('Pharah',
  'Soaring through the air in her combat armor, and armed with a launcher that lays down high-explosive rockets, Pharah is a force to be reckoned with.',
  'Offense',
  'Mobile DPS',
  1
  );

  TRUNCATE game_types CASCADE;

  -- create game_types

  INSERT INTO game_types (name, onAttack)
  VALUES
    ('Assault', true),
    ('Assault', false),
    ('Escort', true),
    ('Escort', false),
    ('Hybrid', true),
    ('Hybrid', false),
    ('Control', true);
