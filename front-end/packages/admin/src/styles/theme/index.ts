

import { createTheme } from '@mui/material/styles';
import { PaletteMode, Theme } from "@mui/material"

export enum Colors {
  blue = "#048ABF",
  green = "#027368",
  red = "#e63946",
  orange = "#b85716",
  grey = "#787878",
  lightgrey = "#E8E8E8",
  violet = "#8404bf"

}

export function GetTheme(mode?: PaletteMode): Theme {
  const theme = createTheme({
    palette: {
      mode: mode ? mode : "light",
      background: {
        default: "#fff"
      },
      primary: {
        light: "#84C9D9",
        main: "#0367A6"
      },
      secondary: {
        main: "#048ABF"
      },
      success: {
        main: "#027368"
      },
      error: {
        main: "#e63946"
      }

    }
  });

  return theme;
}
// #010626
// #0367A6
// #048ABF
// 84C9D9
// #027368
