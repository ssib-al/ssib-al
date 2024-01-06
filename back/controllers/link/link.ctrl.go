package link

import (
	"context"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	"ssib.al/ssib-al-back/config"
	"ssib.al/ssib-al-back/ent/link"
)

var (
	knownSNSBot = []string{
		"facebookexternalhit",
		"twitterbot",
		"googlebot",
		"slack",
		"discord",
		"line",
		"whatsapp",
		"telegram",
		"kakaotalk",
	}
)

func LinkCtrl(c *fiber.Ctx) error {

	domain := c.Hostname()

	uri := c.Params("uri", "")
	if uri == "" {
		return c.Redirect("/")
	}

	dbClient := config.GetDBClient()
	link, err := dbClient.Link.Query().
		Where(link.URIEQ(uri), link.DomainEQ(domain)).
		First(context.Background())
	if err != nil {
		return c.Redirect("/")
	}

	// check user-agent is known sns bot
	if ua := c.Get("User-Agent", ""); ua != "" {
		for _, bot := range knownSNSBot {
			if strings.Contains(strings.ToLower(ua), bot) {
				return c.Redirect(link.TargetURL)
			}
		}
	}

	if _, err = link.QueryOwner().First(context.Background()); err != nil {
		return c.Redirect(link.TargetURL)
	}
	return c.Redirect("/statisticlink/" + link.ID.String() + "/" + url.PathEscape(link.TargetURL))
}
