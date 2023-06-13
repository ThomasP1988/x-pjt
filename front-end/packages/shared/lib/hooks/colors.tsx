import { useContext, Context, createContext, FC, PropsWithChildren } from "react";

export type Color = string;

export type Colors = {
    blue: Color,
    green: Color,
    red: Color,
    orange: Color,
    grey: Color,
    lightgrey: Color,
    violet: Color
}

type Props = {
    colors: Colors
}

const ColorContext = createContext<Colors>({
    blue: "blue",
    green: "green",
    red: "red",
    orange: "orange",
    grey: "grey",
    lightgrey: "lightgrey",
    violet: "violet"
});

export const ColorsProvider: FC<PropsWithChildren<Props>> = ({colors, children}: PropsWithChildren<Props>) => {
    return  <ColorContext.Provider value={colors}>
          {children}
    </ColorContext.Provider>
}

export function useColors(): Colors {
    return useContext(ColorContext);
}