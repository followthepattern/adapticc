export const Icon = ({
  children,
  className,
}: {
  children: any;
  className: string;
}) => {
  return <div className={className}>{children}</div>;
};
