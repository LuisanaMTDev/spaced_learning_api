#!/usr/bin/zsh

(
  source .env
  trap 'cd $HOME/dev/spaced_learning/server' ERR
  cd $HOME/dev/spaced_learning/server/database/sql/schemas && goose sqlite3 $DB_URL_DEV down-to 0 && goose sqlite3 $DB_URL_DEV up && cd $HOME/dev/spaced_learning/server
)
