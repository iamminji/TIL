# Compcation
HBase 의 compaction 관련 코드를 분석하며, 어떻게 동작하는지 알아보고자 한다.
그냥 한줄 한줄 보면서 막 쓰는 거라서 나중에 정리/취합 해야 한다.

compact 관련 코드는 hbase-client, hbase-server 에 각각 있었다.
hbase-client 코드는 Admin proto 에 있는 것으로 그냥 compact 관련해서 RPC 콜을 하는 것으로 보였다.


진또배기는 hbase-server 인 것 같았음

## off peack compaction

## hbase-sever

### StoreFileScanner#getScannersForCompaction

일반 읽기 request 와 충돌을 피하기 위해서 compaction 관련 스캐너를 따로 만든다.
각 스캐너는 storeFile 개수 만큼 만든다.

### Compactor#compact

major 인지 minor 인지에 따라서 dropCache 라는 값을 변경한다.
```
if (request.isMajor() || request.isAllFiles()) {
      dropCache = this.dropCacheMajor;
    } else {
      dropCache = this.dropCacheMinor;
    }
```

잘 이해가 안가는게 `/* Include deletes, unless we are doing a major compaction */`
이런 주석이 있었다.

delete 는 major compaction 때 없어지는데 주석만 보면 major compaction 을 하지 않을때 deletes 를 포함한다고 읽힌다.
앞 뒤로 다시 봐야 할 것 같다.

### Compactor#performCompaction


### 참고
```
org.apache.hadoop.hbase.regionserver.CompactSplitThread.CompactionRunner.doCompaction
         ->CompactSplitThread.selectCompaction //Select the files to be merged based on the policy
         ->HRegion.compact //Execute compact
        ->HStore.compact
            ->DefaultCompactionContext.compact
                ->DefaultCompactor.compact
                                         ->FileDetails fd = getFileDetails(request.getFiles(), request.isAllFiles());//Extract some meta files that need to be merged, because these files are saved as a single hfile
                                         ->scanner = createScanner(store, scanners, scanType, smallestReadPoint, fd.earliestPutTs);//Create a scanner based on these files
                                         ->writer = createTmpWriter(fd, store.throttleCompaction(request.getSize()));//writer to create a new file
                    ->org.apache.hadoop.hbase.regionserver.compactions.Compactor.performCompaction
                        ->do {
                                                         -> hasMore = scanner.next(cells, scannerContext); // traverse data into the file
                            ->CellUtil.setSequenceId(c, 0);
                            ->writer.append(c);
                        ->while (hasMore);
                                         -> writer.appendMetadata(fd.maxSeqId, request.isAllFiles());//write meta
                         ->HStore.moveCompatedFilesIntoPlace //Move to the official directory
                         ->HStore.writeCompactionWalRecord //Write wal
                         ->HStore.replaceStoreFiles //Replace the compact file with the new one. Note that the write lock is added here, that is, the exchange immediately blocks the hstore read and write.
                ->this.lock.writeLock().lock();
                ->this.storeEngine.getStoreFileManager().addCompactionResults(compactedFiles, result);
                ->this.lock.writeLock().unlock();
                         ->HStore.completeCompaction //Close the compactedFile reader, archive history hfile
                ->compactedFile.closeReader(evictOnClose);
                ->this.fs.removeStoreFiles(this.getColumnFamilyName(), compactedFiles);
                ->this.storeSize += r.length();
         ->requestSplit(region)//Check if split is needed
```
- https://programmersought.com/article/52581542959/

