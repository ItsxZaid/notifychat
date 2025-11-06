-- Rename the table from 'campaigns' to 'topics'
ALTER TABLE "campaigns" RENAME TO "topics";

-- Rename the foreign key column in 'channels'
ALTER TABLE "channels" RENAME COLUMN "campaign_id" TO "topic_id";

-- Rename the index on the foreign key
-- (Postgres auto-named this 'channels_campaign_id_idx' from the 001 migration)
ALTER INDEX "channels_campaign_id_idx" RENAME TO "channels_topic_id_idx";

-- Add the new description column to the topics table
-- It's nullable by default, which is good for optional descriptions.
ALTER TABLE "topics" ADD COLUMN "description" text;
