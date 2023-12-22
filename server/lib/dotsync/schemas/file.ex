defmodule Dotsync.Schemas.File do
  use Ecto.Schema
  import Ecto.Changeset

  schema "files" do
    field :path, :string
    field :content, :string
    belongs_to :user, Dotsync.Schemas.User, foreign_key: :user_id

    timestamps()
  end

  @doc false
  def changeset(file, attrs) do
    file
    |> cast(attrs, [:path, :content])
    |> validate_required([:path, :content])
    |> unique_constraint(:unique_file_path, name: :unique_file_path_for_user)
  end
end
