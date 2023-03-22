import clsx from "clsx";
import { HTMLProps } from "react";

interface SelectProps extends Omit<HTMLProps<HTMLSelectElement>, "size"> {
  size?: "xs" | "sm" | "md" | "lg" | "xl";
  variant?: "success" | "danger";
}

const Input: React.FC<SelectProps> = ({
  className,
  size,
  variant,
  ...rest
}) => {
  return (
    <div className="input-control">
      <select
        className={clsx(
          "select",
          { [`text-${variant} input-${variant}`]: variant },
          { [`input--${size}`]: size && size !== "md" },
          className
        )}
        {...rest}
      />
    </div>
  );
};

Input.defaultProps = { type: "text" };
