import clsx from "clsx";
import { HTMLProps, PropsWithChildren } from "react";

interface IModalProps
  extends Omit<PropsWithChildren<HTMLProps<HTMLDivElement>>, "size"> {
  animation?: "dropdown" | "zoom-in" | "zoom-out";
  size?: "small" | "normal" | "large";
}

interface IModal extends React.FC<IModalProps> {
  Overlay: typeof ModalOverlay;
  Content: typeof ModalContent;
  Header: typeof ModalHeader;
  Title: typeof ModalTitle;
  CloseIcon: typeof ModalCloseIcon;
  Body: typeof ModalBody;
  Footer: typeof ModalFooter;
}

const Modal: IModal = ({ className, size, animation, ...rest }) => (
  <div
    className={clsx(
      "modal",
      { [`modal-${size}`]: size && size !== "normal" },
      { [`modal-animated--${animation}`]: animation },
      className
    )}
    {...rest}
  />
);

const ModalOverlay: React.FC<HTMLProps<HTMLAnchorElement>> = ({
  className,
  ...rest
}) => {
  return <a className={clsx("modal-overlay close-btn", className)} {...rest} />;
};

ModalOverlay.defaultProps = { href: "#", "aria-label": "Close" };
ModalOverlay.displayName = "Modal.Overlay";
Modal.Overlay = ModalOverlay;

const ModalContent: React.FC<PropsWithChildren<HTMLProps<HTMLDivElement>>> = ({
  className,
  ...rest
}) => <div className={clsx("modal-content", className)} {...rest} />;

ModalContent.defaultProps = { role: "document" };
ModalContent.displayName = "Modal.Content";
Modal.Content = ModalContent;

const ModalHeader: React.FC<PropsWithChildren<HTMLProps<HTMLDivElement>>> = ({
  className,
  ...rest
}) => <div className={clsx("modal-header", className)} {...rest} />;

ModalHeader.displayName = "Modal.Header";
Modal.Header = ModalHeader;

const ModalTitle: React.FC<PropsWithChildren<HTMLProps<HTMLDivElement>>> = ({
  className,
  ...rest
}) => <div className={clsx("modal-title", className)} {...rest} />;

ModalTitle.displayName = "Modal.Title";
Modal.Title = ModalTitle;

const ModalCloseIcon: React.FC<HTMLProps<HTMLAnchorElement>> = ({
  className,
  ...rest
}) => (
  <a className={clsx("u-pull-right", className)} {...rest}>
    <span className="icon">
      <svg
        aria-hidden="true"
        focusable="false"
        data-prefix="fas"
        data-icon="times"
        className="svg-inline--fa fa-times fa-w-11 fa-wrapper pt-1"
        role="img"
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 352 512"
        height={20}
      >
        <path
          fill="currentColor"
          d="M242.72 256l100.07-100.07c12.28-12.28 12.28-32.19 0-44.48l-22.24-22.24c-12.28-12.28-32.19-12.28-44.48 0L176 189.28 75.93 89.21c-12.28-12.28-32.19-12.28-44.48 0L9.21 111.45c-12.28 12.28-12.28 32.19 0 44.48L109.28 256 9.21 356.07c-12.28 12.28-12.28 32.19 0 44.48l22.24 22.24c12.28 12.28 32.2 12.28 44.48 0L176 322.72l100.07 100.07c12.28 12.28 32.2 12.28 44.48 0l22.24-22.24c12.28-12.28 12.28-32.19 0-44.48L242.72 256z"
        ></path>
      </svg>
    </span>
  </a>
);

ModalCloseIcon.displayName = "Modal.CloseIcon";
ModalCloseIcon.defaultProps = { href: "#", "aria-label": "Close" };
Modal.CloseIcon = ModalCloseIcon;

const ModalBody: React.FC<HTMLProps<HTMLDivElement>> = ({
  className,
  ...rest
}) => <div className={clsx("modal-body", className)} {...rest} />;

ModalBody.displayName = "Modal.Body";
Modal.Body = ModalBody;

const ModalFooter: React.FC<HTMLProps<HTMLDivElement>> = ({
  className,
  ...rest
}) => <div className={clsx("modal-footer", className)} {...rest} />;

ModalFooter.displayName = "Modal.Footer";
Modal.Footer = ModalFooter;

export default Modal;
