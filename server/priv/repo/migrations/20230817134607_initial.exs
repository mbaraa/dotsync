defmodule Dotsync.Repo.Migrations.Initial do
  use Ecto.Migration

  def change do
    create table("users") do
      add :email, :string

      timestamps()
    end

    create unique_index(:users, [:email], name: :unique_email)

    create table("files") do
      add :path, :string
      add :content, :string
      add :user_id, references(:users, on_delete: :delete_all)

      timestamps()
    end

    create unique_index(:files, [:path, :user_id], name: :unique_file_path_for_user)
  end
end
