import { Router } from 'express';
import RequestError from '../lib/request-error';
import * as assetDefinitionsHandler from '../handlers/asset-definitions';
import { constants } from '../lib/utils';
import Ajv from 'ajv';

const ajv = new Ajv();

const router = Router();

router.get('/', async (req, res, next) => {
  try {
    const skip = Number(req.query.skip || 0);
    const limit = Number(req.query.limit || constants.DEFAULT_PAGINATION_LIMIT);
    if (isNaN(skip) || isNaN(limit)) {
      throw new RequestError('Invalid skip / limit', 400);
    }
    res.send(await assetDefinitionsHandler.handleGetAssetDefinitionsRequest(skip, limit));
  } catch (err) {
    next(err);
  }
});

router.get('/:assetDefinitionID', async (req, res, next) => {
  try {
    const assetDefinitionID = Number(req.params.assetDefinitionID);
    if (isNaN(assetDefinitionID)) {
      throw new RequestError('Asset definition ID', 400);
    }
    res.send(await assetDefinitionsHandler.handleGetAssetDefinitionRequest(assetDefinitionID));
  } catch (err) {
    next(err);
  }
});

router.post('/', async (req, res, next) => {
  try {
    if (!req.body.name || req.body.name === '') {
      throw new RequestError('Missing or invalid asset definition name', 400);
    }
    if (!req.body.author) {
      throw new RequestError('Missing asset definition author', 400);
    }
    if (typeof req.body.isContentPrivate !== 'boolean') {
      throw new RequestError('Missing asset definition content privacy', 400);
    }
    if (req.body.descriptionSchema && !ajv.validateSchema(req.body.descriptionSchema)) {
      throw new RequestError('Invalid description schema', 400);
    }
    if (req.body.contentSchema && !ajv.validateSchema(req.body.contentSchema)) {
      throw new RequestError('Invalid content schema', 400);
    }
    await assetDefinitionsHandler.handleCreateAssetDefinitionRequest(req.body.name, req.body.isContentPrivate,
      req.body.author, req.body.descriptionSchema, req.body.contentSchema);
    res.send({ status: 'submitted' });
  } catch (err) {
    next(err);
  }
});

export default router;