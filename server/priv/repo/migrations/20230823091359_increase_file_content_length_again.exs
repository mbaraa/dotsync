defmodule Dotsync.Repo.Migrations.IncreaseFileContentLengthAgain do
  use Ecto.Migration

  def change do
    alter table("files") do
      modify :content, :mediumtext
    end
  end
end
