import clsx from "clsx";
import { HTMLProps, PropsWithChildren } from "react";

type ButtonColor =
  | "transparent"
  | "light"
  | "dark"
  | "black"
  | "primary"
  | "link"
  | "info"
  | "success"
  | "warning"
  | "danger";

type ButtonSize = "xs" | "sm" | "lg" | "xl";

interface ButtonProps
  extends Omit<PropsWithChildren<HTMLProps<HTMLButtonElement>>, "size"> {
  asElement?: React.ElementType;
  color?: ButtonColor;
  outline?: boolean;
  size?: ButtonSize;
  animated?: boolean;
  loading?: boolean;
  loadingSide?: "left" | "right";
  pilled?: boolean;
  circle?: boolean;
  hideText?: boolean;
}

interface IButton extends React.FC<ButtonProps> {
  Group: typeof ButtonGroup;
}

const Button: IButton = ({
  asElement: Component = "button",
  children,
  color,
  size,
  outline,
  animated,
  loading,
  hideText,
  className,
  disabled,
  loadingSide,
  pilled,
  circle,
  ...rest
}) => {
  return (
    <Component
      className={clsx(
        "btn",
        { "btn-animated": animated },

        { [`btn-${color}`]: color },
        { [`btn--${size}`]: size },
        { "btn--pilled": pilled },
        { "btn--circle": circle },
        { "btn--disabled": disabled && Component !== "button" },
        { [`loading-${loadingSide}`]: loadingSide },
        { loading },
        { outline },

        className
      )}
      disabled={disabled}
      {...rest}
    >
      {children}
    </Component>
  );
};

interface ButtonGroupProps
  extends PropsWithChildren<HTMLProps<HTMLDivElement>> {
  groupFill?: boolean;
}

const ButtonGroup: React.FC<ButtonGroupProps> = ({
  className,
  groupFill,
  ...rest
}) => (
  <div
    className={clsx("btn-group", { "btn-group-fill": groupFill }, className)}
    {...rest}
  />
);

Button.Group = ButtonGroup;

export default Button;
