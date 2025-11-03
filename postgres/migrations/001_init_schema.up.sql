CREATE TABLE "campaigns" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "name" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "channels" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "campaign_id" uuid NOT NULL REFERENCES "campaigns"("id") ON DELETE CASCADE,
  "type" text NOT NULL, -- 'telegram', 'whatsapp'
  "config" jsonb NOT NULL,
  "template" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- Add an index to quickly find channels for a campaign
CREATE INDEX ON "channels" ("campaign_id");
