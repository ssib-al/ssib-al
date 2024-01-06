export const domain =
  import.meta.env.VITE_MODE === 'development' ||
  import.meta.env.VITE_MODE === 'staging'
    ? [
        import.meta.env.VITE_DOMAIN_PREFIX + '.' + 'ssib.al',
        import.meta.env.VITE_DOMAIN_PREFIX + '.' + 'zirr.al',
        import.meta.env.VITE_DOMAIN_PREFIX + '.' + 'niae.me',
      ]
    : ['ssib.al', 'zirr.al', 'niae.me'];
