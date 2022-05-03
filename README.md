# stocksheck
An application to download the status of the stocks and shares I own, and allow me to review them

# High Level Design

1. Allow a spreadsheet to be downloaded from a source (Google Drive, Amazon Docs, S3, etc)
2. Read a list of Stocks & Shares
3. Retrieve up to date prices and status for those Stocks & Shares from a source (Yahoo, etc)
4. React to any alerts
5. Generate a useful report and send to a destination (email, Google Drive, Amazon Docs, S3, etc)
6. Update spreadsheet source with last processed date

# Nice to have

1. Show a list of low P/E stocks

# Sources to support

1. Google Drive
2. iCloud Drive
3. OneDrive
4. S3

# Destinations to support

1. Email
2. Any of the above Sources

# Format of Spreadsheet Required

* Column with name, "Stock"
* Column with name, "Alert"
* Column with name, "Today's Price"
