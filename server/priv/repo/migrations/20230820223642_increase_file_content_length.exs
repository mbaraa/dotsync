defmodule Dotsync.Repo.Migrations.IncreaseFileContentLength do
  use Ecto.Migration

  def change do
    alter table("files") do
      modify :content, :text
    end
  end
end
