import { Trans } from "@lingui/macro"
import Box from "@mui/joy/Box"
import Typography from "@mui/joy/Typography"
import OpenInNewIcon from "@mui/icons-material/OpenInNew"
import Link from "@mui/joy/Link"

export default function Footer() {
  return (
    <Box
      sx={{
        gridColumn: { md: "1 / span 2" },
        borderTop: "1px solid",
        borderColor: "divider",
        p: 2,
      }}
    >
      <Typography level="body-sm" textAlign="center">
        <Trans>
          All data used on this site is taken from the{" "}
          <Link
            href="https://arch-bip.ms.gov.pl/pl/rejestry-i-ewidencje/tlumacze-przysiegli/lista-tlumaczy-przysieglych/search.html"
            target="_blank"
          >
            <span style={{ whiteSpace: "nowrap" }}>
              Bulletin of Public information archive <OpenInNewIcon />
            </span>
          </Link>{" "}
          of the Ministry of Justice of the Republic of Poland.
          <br />
          The data is provided "as is" without warranty of any kind for
          informational purposes only.
        </Trans>
      </Typography>

      <Typography level="body-sm" textAlign="center">
        <Trans>Get in touch</Trans>:{" "}
        <Link href="mailto:contact@sworntranslatormap.com">
          contact@sworntranslatormap.com
        </Link>
      </Typography>
    </Box>
  )
}
