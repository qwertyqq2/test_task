
## Test task - parsing data from tender sites

The task is to parse a website containing
information about tenders. In my case, the site was selected https://www.rts-tender.ru . At the same time , the same problem was solved for https://www.etp-ets.ru, as a test

### Installation

    git clone github.com/qwertyqq2/test_task
    
    cd test_task


### Usage

    ./main -rts=true

or

    ./main -etp=true

### Inside

Since different pages display their data differently, I decided to create a common interface for all parsers.

    // Req is an interface that starts and cancels the process of searching for data from a web page.
    type Req interface {

        //SendRequest sends a request to a web page and returns
        // a channel through which we will receive data from it.
        SendRequest(ctx context.Context) <-chan *data.Data

        //Cancel completes the process of receiving data, while closing the channel
        Cancel()

        //are we in the process of getting the data?
        Proccessing() bool
    }



### Example

etp := req.NewEtp()
ch := etp.SendRequest(ctx)
for {
	select {
	    case <-ctx.Done():
			return

		case d := <-ch:
			fmt.Println(d.String())
	}
}

