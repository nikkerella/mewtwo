An example to catch an error in goroutines and cancel the other goroutines without errgroup.
- Use a dedicated channel to signal the error to the main function.
- The main function cancel the context, then propagate to other goroutines.