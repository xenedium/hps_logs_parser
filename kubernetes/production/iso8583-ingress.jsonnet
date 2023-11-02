local host = std.extVar("host");

[
  {
    apiVersion: 'networking.k8s.io/v1',
    kind: 'Ingress',
    metadata: {
      name: 'iso8583backend-ingress',
      annotations: {
        'cert-manager.io/issuer': 'letsencrypt-prod',
        'nginx.ingress.kubernetes.io/proxy-body-size': '0',
        'nginx.ingress.kubernetes.io/proxy-read-timeout': '600',
        'nginx.ingress.kubernetes.io/proxy-send-timeout': '600',
      },
    },
    spec: {
      ingressClassName: 'nginx',
      tls: [
        {
          hosts: [
            host,
          ],
          secretName: 'iso8583-tls',
        },
      ],
      rules: [
        {
          host: host,
          http: {
            paths: [
              {
                path: '/api',
                pathType: 'Prefix',
                backend: {
                  service: {
                    name: 'iso8583backend',
                    port: {
                      number: 8000,
                    },
                  },
                },
              },
            ],
          },
        },
      ],
    },
  },
  {
    apiVersion: 'networking.k8s.io/v1',
    kind: 'Ingress',
    metadata: {
      name: 'iso8583frontend-ingress',
      annotations: {
        'cert-manager.io/issuer': 'letsencrypt-prod',
      },
    },
    spec: {
      ingressClassName: 'nginx',
      tls: [
        {
          hosts: [
            host,
          ],
          secretName: 'iso8583-tls',
        },
      ],
      rules: [
        {
          host: host,
          http: {
            paths: [
              {
                path: '/',
                pathType: 'Prefix',
                backend: {
                  service: {
                    name: 'iso8583frontend',
                    port: {
                      number: 80,
                    },
                  },
                },
              },
            ],
          },
        },
      ],
    },
  },
]