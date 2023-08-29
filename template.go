package echoroutesview

const routerViewer = `
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Routes</title>

    <link href="https://cdn.jsdelivr.net/npm/daisyui@3.6.2/dist/full.css" rel="stylesheet" type="text/css" />
    <script defer src="https://cdn.tailwindcss.com"></script>
    <style>
        main {
            width: 60rem;
            margin: 0 auto;
        }
    </style>

</head>

<body>
    <main>
        <div class="stats shadow text-primary-content w-full">
            <div class="stat">
              <div class="stat-title">Total Registered Routes</div>
              <div class="stat-value">{{ len . }}</div>
            </div>

          </div>

          <div class="overflow-x-auto">
            <table class="table table-zebra">
              <thead>
                <tr>
                  <th></th>
                  <th>Path</th>
                  <th>Handler Method</th>
                </tr>
              </thead>
              <tbody>
                {{ range $index, $route := . }}
                    <tr>
                        <th>{{ add $index  1 }}</th>
                        <td>{{ $route.Path }}
                            {{ if (eq $route.Method "GET") }}
                                <div class="badge badge-success">{{ $route.Method }}</div>
                            {{ else if (eq $route.Method "POST")}}
                                <div class="badge badge-info">{{ $route.Method }}</div>
                            {{ else if (eq $route.Method "PUT")}}
                                <div class="badge badge-warning">{{ $route.Method }}</div>
                            {{ else if (eq $route.Method "DELETE")}}
                                <div class="badge badge-error">{{ $route.Method }}</div>
                            {{ else if (eq $route.Method "PATCH")}}
                                <div class="badge badge-primary">{{ $route.Method }}</div>
                            {{ else }}
                                <div class="badge badge-neutral">{{ $route.Method }}</div>
                            {{ end }}

                        </td>
                        <td>{{ $route.Name }}</td>
                    </tr>
                {{ end }}

              </tbody>
            </table>
          </div>
    </main>

</body>

</html>
`
