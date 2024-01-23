



if: github.event_name == 'pull_request' && github.event.action == 'closed' && github.event.pull_request.merged == true && github.ref == 'refs/heads/master'
