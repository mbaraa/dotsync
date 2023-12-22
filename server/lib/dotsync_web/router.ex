defmodule DotsyncWeb.Router do
  use DotsyncWeb, :router

  pipeline :api do
    plug :accepts, ["json"]
  end

  scope "/user", DotsyncWeb do
    pipe_through :api

    post "/login", User, :login
    post "/login/verify", User, :verify_login
    delete "/", User, :delete_user
  end

  scope "/file", DotsyncWeb do
    pipe_through :api

    post "/", File, :add_file
    post "/add-directory", File, :add_directory
    delete "/", File, :delete_file
    get "/", File, :get_synced_files
    get "/download", File, :download_files
    post "/upload", File, :upload_files
  end
end
