import clsx from "clsx";
import { HTMLProps, useId } from "react";

interface InputProps extends Omit<HTMLProps<HTMLInputElement>, "size"> {
  size?: "xs" | "sm" | "md" | "lg" | "xl";
}

const Input: React.FC<InputProps> = ({
  className,
  size,
  label,
  id,
  ...rest
}) => {
  const generatedID = useId();
  const inputID = id ?? generatedID;

  return (
    <div className="input-control">
      {label && <label htmlFor={inputID}>{label}</label>}
      <input
        id={inputID}
        className={clsx(
          { [`input--${size}`]: size && size !== "md" },
          className
        )}
        {...rest}
      />
    </div>
  );
};

Input.defaultProps = { type: "text" };

export default Input;
