import { SHA } from '../RepositoryData';
import { IElementDiff, IBlobDiff } from './Diff';

export type IDatasetBlob = IPathDatasetBlob | IS3DatasetBlob;

export interface IPathDatasetBlob {
  type: 'path';
  category: 'dataset';
  components: IPathDatasetComponentBlob[];
}

export interface IS3DatasetBlob {
  category: 'dataset';
  type: 's3';
  components: IS3DatasetComponentBlob[];
}
export interface IS3DatasetComponentBlob {
  path: IPathDatasetComponentBlob;
}

export interface IPathDatasetComponentBlob {
  path: string; // Full path to the file
  size: number;
  lastModifiedAtSource: Date;
  sha256: SHA;
  md5: string;
}

export const datasetBlobTypes: Record<
  IDatasetBlob['type'],
  IDatasetBlob['type']
> = {
  s3: 's3',
  path: 'path',
};

// diff

export type IPathDatasetBlobDiff = IBlobDiff<
  IPathDatasetBlobDiffData,
  IPathDatasetBlob['category'],
  IPathDatasetBlob['type']
>;
export type IPathDatasetBlobDiffData = {
  type: 'path';
  category: 'dataset';
  components: IPathDatasetComponentBlobDiff[];
};

export type IS3DatasetBlobDiff = IBlobDiff<
  IS3DatasetBlobDiffData,
  IS3DatasetBlob['category'],
  IS3DatasetBlob['type']
>;
export type IS3DatasetBlobDiffData = {
  category: IS3DatasetBlob['category'];
  type: IS3DatasetBlob['type'];
  components: Array<{
    path: IPathDatasetComponentBlobDiff;
  }>;
};

export type IPathDatasetComponentBlobDiff = IElementDiff<
  IPathDatasetComponentBlob
>;

export type IDatasetBlobDiff = IPathDatasetBlobDiff | IS3DatasetBlobDiff;
export type IDatasetBlobDiffData =
  | IPathDatasetBlobDiffData
  | IS3DatasetBlobDiffData;
