defmodule Dotsync.Mailer do
  use Swoosh.Mailer, otp_app: :dotsync
  import Swoosh.Email

  @typedoc """
    A valid email address, and login token.
  """
  @type token :: charlist()
  @type email :: charlist()
  def send_login_token(token, email) do
    new()
    |> to(email)
    |> from({"Dotsync", "pub@mbaraa.com"})
    |> subject("Login token")
    |> html_body(get_login_token_html(token))
    |> text_body("Hello, here's your login token: #{token}, DON'T SHARE IT WITH ANYONE!")
    |> Dotsync.Mailer.deliver()
  end

  defp get_login_token_html(token) do
    """
    <html lang="en">
    <head>
    <style>
      @import url("https://fonts.googleapis.com/css2?family=IBM+Plex+Sans+Condensed&display=swap");
    </style>
    </head>
    <body>
    <div
      style="
        font-family: &quot;IBM Plex Sans Condensed&quot;, sans-serif;
        min-width: 500px;
        overflow: auto;
        line-height: 2;
      "
    >
      <div style="margin: 50px auto; width: 70%; padding: 20px 0">
        <div style="border-bottom: 1px solid #eee">
          <a
            href="https://dotsync.org"
            style="
              font-size: 1.4em;
              color: #00466a;
              text-decoration: none;
              font-weight: 600;
            "
            >Dotsync</a
          >
        </div>
        <p style="font-size: 1.1em">Hi,</p>
        <p>
          Thanks for using Dotsync, keep in mind that your email and data,
          <strong>are not</strong> shared or viewed by anyone! <br/>And below lies
          your login token, don't share it with anyone, in order to keep your
          account safe 😁
        </p>
        <h2
          style="
            background: #00466a;
            margin: 0 auto;
            width: max-content;
            padding: 0 10px;
            color: white;
            border-radius: 4px;
            font-size: 12px;
            max-width: 500px;
            word-break: break-word;
          "
        >
          #{token}
        </h2>
        <p style="font-size: 0.9em">
          Regards,<br />Baraa Al-Masri :: Dotsync's Admin
        </p>

        <hr style="border: none; border-top: 1px solid #eee" />
        <div
          style="
            float: right;
            padding: 8px 0;
            color: #aaa;
            font-size: 0.8em;
            line-height: 1;
            font-weight: 300;
          "
        >
          <p>Dotsync</p>
          <p>
            <a href="mailto:pub@mbaraa.com">pub@mbaraa.com</a>
            <br />
            <a href="https://mbaraa.com">mbaraa.com</a>
            <br />
            <a href="https://dotsync.org">dotsync.org</a>
          </p>
        </div>
      </div>
    </div>
    </body>
    </html>
    """
  end
end
