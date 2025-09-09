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
                                <div class="badge badge-success">{{ $route.Method }}<a href="{{ $route.Path }}" target="_blank" class="inline"><svg width="20px" height="20px" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg"><g ><path d="M10.0002 5H8.2002C7.08009 5 6.51962 5 6.0918 5.21799C5.71547 5.40973 5.40973 5.71547 5.21799 6.0918C5 6.51962 5 7.08009 5 8.2002V15.8002C5 16.9203 5 17.4801 5.21799 17.9079C5.40973 18.2842 5.71547 18.5905 6.0918 18.7822C6.5192 19 7.07899 19 8.19691 19H15.8031C16.921 19 17.48 19 17.9074 18.7822C18.2837 18.5905 18.5905 18.2839 18.7822 17.9076C19 17.4802 19 16.921 19 15.8031V14M20 9V4M20 4H15M20 4L13 11" stroke="#000000" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></g></svg></a></div>
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
