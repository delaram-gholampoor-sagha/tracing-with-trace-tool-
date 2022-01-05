# tracing-with-trace-tool-
Tracing Go code

##### tracing allows you to pass context through your system and evaluate where you are being held up , whether its by a third-party API call , a slow messaging quue , or an O(n2) function . tracing will help you figure out where the bottleneck resides

##### the go tracing tool hooks into the goroutine scheduler so that it can produce meaningful information about goroutines.

